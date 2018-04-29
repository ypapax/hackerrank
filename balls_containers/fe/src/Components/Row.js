import React, { Component } from 'react';
class Row extends Component {
  render() {
        if (!this.props.cells) {
            return ""
        }
        let cells = this.props.cells.map((v, i) => {
           return (<td key={i} className="cell">{v}</td>)
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
