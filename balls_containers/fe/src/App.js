import React, {Component} from 'react';
import './App.css';
import Arranger from "./Components/Arranger";
import $ from "jquery";

const backend = "http://localhost:8083";

class App extends Component {
    constructor() {
        super()
        this.state = {
            matrices: []
        }
    }

    showError(msg) {
        alert(msg);
    }

    matrixChanged(m, cb){
        debugger;
        let mm = this.state.matrices;
        if (!mm || !mm.length) {
            mm = [m];
        } else {
            mm.push(m);
        }

        this.setState({
            "matrices": mm
        }, function () {
            console.info("new state", this.state);
            if (cb) {
                cb();
            }
            window.scrollTo(0,document.body.scrollHeight); // scroll to bottom
        })
    }

    parse(input) {
        $.ajax({
            type: "POST",
            url: backend + "/api/v1/parse",
            data: JSON.stringify({Params: [input]}),
            success: function (data) {
                console.info("data", data);
                if (this.hasError(data)) {
                    return;
                }
                this.matrixChanged(data[0]/*, this.arrange.bind(this)*/);
            }.bind(this),
            dataType: "json",
            error: function (e) {
                console.error(e);
            }
        });
    }

    lastMatrix(){
        return this.state.matrices[this.state.matrices.length - 1];
    }

    arrange() {
        $.ajax({
            type: "POST",
            url: backend + "/api/v1/arrange",
            data: JSON.stringify({Params: [JSON.stringify(this.lastMatrix())]}),
            success: function (data) {
                console.info("data", data);
                if (this.hasError(data)) {
                    return;
                }
                this.matrixChanged(data.m);
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
                    JSON.stringify(this.lastMatrix()),
                    senderRow.toString(), senderColumn.toString(),
                    targetRow.toString(), targetColumn.toString()
                ]
            }),
            success: function (data) {
                console.info("data", data);
                if (this.hasError(data)) {
                    return;
                }
                this.matrixChanged(data);
            }.bind(this),
            dataType: "json",
            error: function (e) {
                console.error(e);
            }
        });
    }

    hasError(data){
        if (data.hasOwnProperty("reason")) {
            this.showError(data.reason);
            return data;
        }
        return null;
    }


    inputChanged(input) {
        console.info("inputChanged", input);
        this.parse(input);
    }

    onDrop(senderRow, senderColumn, targetRow, targetColumn) {
        this.swap(senderRow, senderColumn, targetRow, targetColumn);
    }

    render() {
        return (
            <div className="App">
                <Arranger
                    matrices={this.state.matrices}
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
