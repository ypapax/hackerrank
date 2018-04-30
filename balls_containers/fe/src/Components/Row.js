import React, {Component} from 'react';
import Cell from "./Cell";

class Row extends Component {

    render() {
        if (!this.props.cells) {
            return ""
        }
        let cells = this.props.cells.map((v, i) => {
            let matrix=this.props.matrix;
            let row=this.props.row;
            let column = i;
            let key=matrix.toString() + "_" + row.toString() + "_" + column.toString();
            return (<Cell
                value={v}
                column={column}
                row={this.props.row}
                onDrop={this.props.onDrop.bind(this)}
                key={key}
                onDragStart={this.props.onDragStart}
                onDragOver={this.props.onDragOver}
                onDragLeave={this.props.onDragLeave.bind(this)}
            />)
        });
        console.info("cells", cells);
        return (
            <tr>
                {cells}
            </tr>
        );
    }
}

export default Row;
