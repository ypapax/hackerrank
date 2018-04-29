import React, { Component } from 'react';
import Cell from "./Cell";
class Row extends Component {
  render() {
        if (!this.props.cells) {
            return ""
        }
        let cells = this.props.cells.map((v, i) => {
           return (<Cell value={v} key={i}/>)
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
