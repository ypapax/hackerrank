import React, {Component} from 'react';

class Cell extends Component {
    allowDrop(e) {
        e.preventDefault();
    }

    onDrop(ev) {
        ev.preventDefault();
        var senderRow = ev.dataTransfer.getData("row");
        var senderCol = ev.dataTransfer.getData("column");
        this.props.onDrop(senderRow, senderCol, this.props.row, this.props.column);
    }

    dragStart(ev) {
        ev.dataTransfer.setData("row", this.props.row);
        ev.dataTransfer.setData("column", this.props.column);
    }

    render() {
        return (
            <td className="cell" draggable="true" onDrop={this.onDrop.bind(this)}
                onDragStart={this.dragStart.bind(this)}
                onDragOver={this.allowDrop.bind(this)}
            >
                {this.props.value}
            </td>
        )
    }
}

export default Cell;
