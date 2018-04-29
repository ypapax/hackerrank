import React, {Component} from 'react';
import './App.css';
import Arranger from "./Components/Arranger";
import $ from "jquery";

const backend = "http://localhost:8083";

class App extends Component {
    constructor() {
        super()
        this.state = {}
    }

    parse(input) {
        $.ajax({
            type: "POST",
            url: backend + "/api/v1/parse",
            data: JSON.stringify({Params: [input]}),
            success: function (data) {
                console.info("data", data);
                this.setState({
                    "matrix": data[0]
                }, function () {
                    console.info("new state", this.state);
                    this.arrange();
                })
            }.bind(this),
            dataType: "json",
            error: function (e) {
                console.error(e);
            }
        });
    }

    arrange() {
        $.ajax({
            type: "POST",
            url: backend + "/api/v1/arrange",
            data: JSON.stringify({Params: [JSON.stringify(this.state.matrix)]}),
            success: function (data) {
                console.info("data", data);
                this.setState({
                    "matrix": data.m
                }, function () {
                    console.info("new state", this.state);
                })
            }.bind(this),
            dataType: "json",
            error: function (e) {
                console.error(e);
            }
        });
    }

    inputChanged(input) {
        console.info("inputChanged", input);
        this.parse(input);
    }

    onDrop(row, column) {
        let m = this.state.matrix;
        debugger;
        m[row][column] = 666;
        this.setState({
            "matrix": m
        })
    }

    render() {
        return (
            <div className="App">
                <Arranger
                    matrix={this.state.matrix}
                    onParse={this.parse.bind(this)}
                    onArrange={this.arrange.bind(this)}
                    onInputChanged={this.inputChanged.bind(this)}
                    onDrop={this.onDrop.bind(this)}
                />
            </div>
        );
    }
}

export default App;
