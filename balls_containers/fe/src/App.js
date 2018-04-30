import React, {Component} from 'react';
import './App.css';
import Arranger from "./Components/Arranger";
import $ from "jquery";

const backend = "http://localhost:8083";

class App extends Component {
    constructor() {
        super()
        this.state = {
            matrices: [],
            info: []
        }
    }

    showError(msg) {
        alert(msg);
    }

    matrixChanged(m, info, cb) {
        let mm = this.state.matrices;
        if (!mm || !mm.length) {
            mm = [m];
        } else {
            mm.push(m);
        }

        let infos = this.state.info;
        infos.push(info)

        this.setState({
            "matrices": mm,
            "info": infos
        }, function () {
            console.info("new state", this.state);
            if (cb) {
                cb();
            }
            window.scrollTo(0, document.body.scrollHeight); // scroll to bottom
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
                this.matrixChanged(data[0], "parsing", this.arrange.bind(this));
            }.bind(this),
            dataType: "json",
            error: function (e) {
                console.error(e);
            }
        });
    }

    lastMatrix() {
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
                this.matrixChanged(data.m, "arranging");
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
                let amount = data.swap.Amount;
                let msg1 = "moved " + amount + " balls of type " + senderColumn + " from box " + senderRow + " to box " + targetRow;
                let msg2 = "moved " + amount + " balls of type " + targetColumn + " from box " + targetRow + " to box " + senderRow;
                let msg = msg1 + "<br/>" + msg2;
                this.matrixChanged(data.matrix, msg);
            }.bind(this),
            dataType: "json",
            error: function (e) {
                console.error(e);
            }
        });
    }

    hasError(data) {
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
                    info={this.state.info}
                />
            </div>
        );
    }
}

export default App;
