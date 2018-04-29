import React, { Component } from 'react';
class Row extends Component {
    constructor(props){
        super(props);
    }

  render() {
        if (!this.props.cells) {
            return ""
        }
        let cells = this.props.cells.map(v => {
           return (<span className="cell">{v}</span>)
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
