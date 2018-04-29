import React, {Component} from 'react';
import Cell from "./Cell";

class Row extends Component {
    onDrop(column) {
        let row = this.props.row;
        this.props.onDrop(row, column);
    }

    render() {
        if (!this.props.cells) {
            return ""
        }
        let cells = this.props.cells.map((v, i) => {
            return (<Cell
                value={v}
                column={i}
                onDrop={this.onDrop.bind(this)}
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
