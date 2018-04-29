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

    showError(msg) {
        alert(msg);
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

    swap(senderRow, senderColumn, targetRow, targetColumn) {
        $.ajax({
            type: "POST",
            url: backend + "/api/v1/swap",
            data: JSON.stringify({
                Params: [
                    JSON.stringify(this.state.matrix),
                    senderRow.toString(), senderColumn.toString(),
                    targetRow.toString(), targetColumn.toString()
                ]
            }),
            success: function (data) {
                console.info("data", data);
                if (data.hasOwnProperty("reason")) {
                    this.showError(data.reason);
                    return;
                }
                this.setState({
                    "matrix": data
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

    onDrop(senderRow, senderColumn, targetRow, targetColumn) {
        let m = this.state.matrix;
        this.swap(senderRow, senderColumn, targetRow, targetColumn);
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
