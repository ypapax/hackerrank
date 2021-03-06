package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"os"

	"github.com/golang/glog"
)

func arrangeAPI(bindServer string) {
	apiFuncs := initCmds()
	http.HandleFunc("/api/v1/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		log.Printf("URL path: %s", r.URL.Path[1:])

		fields := strings.Split(r.URL.Path[1:], "/")
		if len(fields) < 3 {
			log.Printf("invalid endpoint path %s", r.URL.Path[1:])
			http.Error(w, "invalid endpoint", http.StatusBadRequest)
			return
		}
		ep := fields[2]
		apiFunc, ok := apiFuncs[ep]
		if !ok {
			log.Printf("invalid endpoint %s", ep)
			http.Error(w, "invalid command", http.StatusBadRequest)
			return
		}

		var req = &Request{UriTunnel: fields[2:]}
		if r.Method == "POST" {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("%+v", err)
				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}
			log.Printf("Retrieve body %s", string(b))
			if err := json.Unmarshal(b, &req); err != nil {
				log.Printf("can't decode request msg")
				http.Error(w, "can't decode body", http.StatusBadRequest)
				return
			}
		}
		cmd := apiCmd{
			RemoteAddr: r.RemoteAddr, Cmd: req, Req: r}
		log.Printf("ep %+v", ep)
		resp, err := apiFunc.Handler(cmd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendResp(w, resp)
	})
	log.Printf("listening %+v", bindServer)
	log.Fatal(http.ListenAndServe(bindServer, nil))
}

type apiFunc struct {
	Handler handler
}

func sendResp(w http.ResponseWriter, resp *Response) {
	result := resp.Result
	if resp.Error != nil {
		result = resp.Error
	}
	payload, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "resource error", http.StatusInternalServerError)
		return
	}
	glog.V(6).Infof("Payload: %s", string(payload))
	w.Write(payload)
}

func initCmds() map[string]*apiFunc {
	return map[string]*apiFunc{
		"arrange": {Handler: arrangeHandler},
		"parse":   {Handler: parseHandler},
		"swap":    {Handler: swapHandler},
	}
}

type apiCmd struct {
	RemoteAddr string
	Cmd        *Request
	Req        *http.Request
}

type Request struct {
	Directive string
	UriTunnel []string
	Params    []string
}

type Response struct {
	Error  *ErrorResp
	Result interface{}
}

type ErrorResp struct {
	Reason string `json:"reason",omitempty`
	Code   string `json:"code"`
}

type handler func(af apiCmd) (*Response, error)

func arrangeHandler(af apiCmd) (*Response, error) {
	if len(af.Cmd.Params) == 0 {
		err := fmt.Errorf("not enough arguments")
		glog.Info(err)
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	mStr := af.Cmd.Params[0]
	var m [][]int
	err := json.Unmarshal([]byte(mStr), &m)
	if err != nil {
		err = fmt.Errorf("Unable to parse matrix: %+v", err)
		glog.Error(err)
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	isArranged := isPossibleToArrange(m)
	return &Response{Result: map[string]interface{}{"m": m, "isArranged": isArranged}}, nil
}

func parseHandler(af apiCmd) (*Response, error) {
	if len(af.Cmd.Params) == 0 {
		err := fmt.Errorf("not enough arguments")
		glog.Info(err)
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	b := bytes.NewBufferString(af.Cmd.Params[0])
	m, err := readMatrices(b)
	if err != nil {
		err := fmt.Errorf("unable to parse matrix")
		glog.Error(err)
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	return &Response{Result: m}, nil
}

func swapHandler(af apiCmd) (*Response, error) {
	if len(af.Cmd.Params) < 5 {
		err := fmt.Errorf("not enough arguments")
		glog.Info(err)
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	var m [][]int
	err := json.Unmarshal([]byte(af.Cmd.Params[0]), &m)
	if err != nil {
		err := fmt.Errorf("could not parse matrix %+v", err)
		glog.Error(err)
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	boxNumber, err := strconv.Atoi(af.Cmd.Params[1])
	if err != nil {
		err := fmt.Errorf("could not parse int %+v", err)
		glog.Error(err)
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	ballType, err := strconv.Atoi(af.Cmd.Params[2])
	if err != nil {
		err := fmt.Errorf("could not parse int %+v", err)
		glog.Error(err)
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	targetBoxNumber, err := strconv.Atoi(af.Cmd.Params[3])
	if err != nil {
		err := fmt.Errorf("could not parse int %+v", err)
		glog.Error(err)
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	targetBallType, err := strconv.Atoi(af.Cmd.Params[4])
	if err != nil {
		err := fmt.Errorf("could not parse int %+v", err)
		glog.Error(err)
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}

	sw, err := swapByBoxFromToAndBallNumber(boxNumber, ballType, targetBoxNumber, targetBallType, m)
	if err != nil {
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	return &Response{Result: map[string]interface{}{"matrix": m, "swap": sw}}, nil
}

func swapByBoxFromToAndBallNumber(boxNumber, ballTypeNumber, targetBox, ballTypeNumber2 int, m [][]int) (*swapping, error) {
	sw := newSwapping(boxNumber, ballTypeNumber, targetBox, ballTypeNumber2)
	a1 := m[sw.ballMove1.from.row][sw.ballMove1.from.column]
	a2 := m[sw.ballMove2.from.row][sw.ballMove2.from.column]
	log.Println("box", boxNumber, "contains", a1, "balls of type", ballTypeNumber)
	log.Println("box", targetBox, "contains", a2, "balls of type", ballTypeNumber2)
	sw.Amount = int(math.Min(float64(a1), float64(a2)))
	pr := func(msg string) {
		log.Println(msg)
		printMatrix(m, sw.ballMove1.from, sw.ballMove1.to, sw.ballMove2.from, sw.ballMove2.to)
	}

	if sw.Amount == 0 {
		pr("swap Amount is 0")
		return nil, fmt.Errorf("swap Amount is 0")
	}
	pr(fmt.Sprintf("before swap %+v", sw))
	m = swap(m, sw)
	pr("after swap")
	return &sw, nil
}

func swap(m [][]int, sw swapping) [][]int {
	for _, bm := range []ballMove{sw.ballMove1, sw.ballMove2} {
		m[bm.from.row][bm.from.column] -= sw.Amount
		m[bm.to.row][bm.to.column] += sw.Amount
		log.Printf("%+v balls of type %+v had moved from box %+v to box %+v\n", sw.Amount, bm.from.column, bm.from.row, bm.to.row)
	}
	return m
}

func newSwapping(box1, ballType1, box2, ballType2 int) swapping {
	return swapping{
		ballMove1: ballMove{
			point{box1, ballType1, red}, point{box2, ballType1, green},
		},
		ballMove2: ballMove{
			point{box2, ballType2, redBG}, point{box1, ballType2, greenBG},
		},
		Amount: 1,
	}
}

type ballMove struct {
	from, to point
}

type swapping struct {
	ballMove1, ballMove2 ballMove
	Amount               int
}

type point struct {
	row, column int
	color       color
}

type color int

const (
	red     color = 31
	green   color = 32
	redBG   color = 101
	greenBG color = 102
)

func getColor(c color) string {
	if c == 0 {
		return "\033[0m" // turn off the color
	}
	return fmt.Sprintf("\033[0;%+vm", c) // https://misc.flogisoft.com/bash/tip_colors_and_formatting
}

func printMatrix(m [][]int, highlight ...point) {
	maxLength := func() int {
		var max int
		for _, r := range m {
			for _, v := range r {
				if l := len(fmt.Sprintf("%+v", v)); l > max {
					max = l
				}
			}
		}
		return max
	}()

	line := func() {
		l := "-"
		size := maxLength * len(m)
		for i := 0; i <= size; i++ {
			l += "-"
		}
		fmt.Fprint(os.Stderr, l, "\n")
	}
	line()
	for i, r := range m {
		for j, v := range r {
			_ = i
			_ = j
			var clr, noClr string
			for _, h := range highlight {
				if i == h.row && j == h.column {
					clr = getColor(h.color)
					noClr = getColor(0)
					break
				}

			}
			tab := " "
			delta := maxLength + 1 - len(fmt.Sprintf("%+v", v))
			for i := 0; i < delta; i++ {
				tab += " "
			}
			fmt.Fprint(os.Stderr, clr, v, noClr, tab)
		}
		fmt.Fprintln(os.Stderr)
	}
	line()
}
