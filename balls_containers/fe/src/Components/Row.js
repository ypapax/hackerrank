import React, {Component} from 'react';
import Cell from "./Cell";

class Row extends Component {

    render() {
        if (!this.props.cells) {
            return ""
        }
        let cells = this.props.cells.map((v, i) => {
            return (<Cell
                value={v}
                column={i}
                row={this.props.row}
                onDrop={this.props.onDrop.bind(this)}
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
