import React, { Component } from 'react';
class Row extends Component {
  render() {
        if (!this.props.cells) {
            return ""
        }
        let cells = this.props.cells.map((v, i) => {
           return (<span key={i} className="cell">{v}</span>)
        });
        console.info("cells", cells);
    return (
      <div className="row">
          {cells}
      </div>
    );
  }
}

export default Row;
