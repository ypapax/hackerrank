import React, {Component} from 'react';

class Cell extends Component {
    constructor() {
        super();
        this.state = {};
    }

    dragOver(ev) {
        ev.preventDefault();
        this.props.onDragOver(this.props.row, this.props.column);
    }

    onDrop(ev) {
        ev.preventDefault();
        var senderRow = ev.dataTransfer.getData("row");
        var senderCol = ev.dataTransfer.getData("column");
        this.props.onDrop(senderRow, senderCol, this.props.row, this.props.column);
        this.setState({
            dropping: true
        })
    }

    dragStart(ev) {
        ev.dataTransfer.setData("row", this.props.row);
        ev.dataTransfer.setData("column", this.props.column);
        this.setState({
            dragging: true
        })
        this.props.onDragStart(this.props.row, this.props.column);
    }

    render() {
        let dragDropStyle = "";
        if (this.state.dragging) {
            dragDropStyle = "drag";
        } else if (this.state.dropping)  {
            dragDropStyle = "drop";
        }
        let className = "cell " + dragDropStyle;
        return (
            <td className={className} draggable="true" onDrop={this.onDrop.bind(this)}
                onDragStart={this.dragStart.bind(this)}
                onDragOver={this.dragOver.bind(this)}
                onDragLeave={this.props.onDragLeave.bind(this)}
            >
                {this.props.value}
            </td>
        )
    }
}

export default Cell;
