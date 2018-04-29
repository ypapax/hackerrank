package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

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

	m, isArranged := arrangeMatrix(m, false)
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

	if err := swapByBoxFromToAndBallNumber(boxNumber, ballType, targetBoxNumber, targetBallType, m); err != nil {
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	return &Response{Result: m}, nil
}
