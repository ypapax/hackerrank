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

    swap(senderRow, senderColumn, targetRow, targetColumn, cb) {
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
            success: cb,
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
        this.cancelHover(function () {
            if (senderRow == targetRow) {
                console.warn("should be different lines, instead they are both: ", senderRow);
                return;
            }
            if (senderColumn == targetColumn) {
                console.warn("should be different columns, instead they are both: ", senderColumn);
                return;
            }
            this.swap(senderRow, senderColumn, targetRow, targetColumn, function (data) {
                if (this.hasError(data)) {
                    return;
                }
                let amount = data.swap.Amount;
                let msg1 = "moved " + amount + " balls of type " + senderColumn + " from box " + senderRow + " to box " + targetRow;
                let msg2 = "moved " + amount + " balls of type " + targetColumn + " from box " + targetRow + " to box " + senderRow;
                let msg = msg1 + "<br/>" + msg2;
                this.matrixChanged(data.matrix, msg);
            }.bind(this));
        }.bind(this));
    }

    onDragStart(row, column) {
        this.setState({dragging: [row, column]}, function () {
            console.info("new state after drag start", this.state);
        });
    }

    arraysEqual(arr1, arr2) {
        if (arr1.length !== arr2.length)
            return false;
        for (var i = arr1.length; i--;) {
            if (arr1[i] !== arr2[i])
                return false;
        }

        return true;
    }

    // Over is in sense of hover.
    onDragOver(targetRow, targetColumn) {
        let senderRow = this.state.dragging[0];
        let senderColumn = this.state.dragging[1];
        let swapParams = [senderRow, senderColumn, targetRow, targetColumn];
        if (this.state.lastDragOverRequest && this.arraysEqual(this.state.lastDragOverRequest, swapParams)) {
            console.info("this.state.lastDragOverRequest  is the same, skipping", JSON.stringify(swapParams));
            return;
        }
        this.state.lastDragOverRequest = swapParams;
        this.swap(senderRow, senderColumn, targetRow, targetColumn, function (data) {
            if (data.hasOwnProperty("reason")) {
                console.error(data.reason);
                return data;
            }
            let m = data.matrix;
            let mm = this.state.matrices;
            let lastMatrix = mm[mm.length - 1];
            mm[mm.length - 1] = m;
            this.setState({
                lastMatrix: lastMatrix,
                matrices: mm
            })
        }.bind(this));
    }

    cancelHover(cb) {
        if (!this.state.lastMatrix) {
            cb();
            return;
        }
        let mm = this.state.matrices;
        mm[mm.length - 1] = this.state.lastMatrix;

        this.setState({matrices: mm}, function () {
            cb();
        }.bind(this));
    }

    onDragLeave() {
        this.cancelHover(function () {
            console.info("cancelHover in onDragLeave");
        }.bind(this));
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
                    onDragStart={this.onDragStart.bind(this)}
                    onDragOver={this.onDragOver.bind(this)}
                    onDragLeave={this.onDragLeave.bind(this)}

                />
            </div>
        );
    }
}

export default App;
