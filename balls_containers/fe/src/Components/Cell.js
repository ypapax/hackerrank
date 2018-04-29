import React, { Component } from 'react';
class Cell extends Component {
  render() {
        return (<td className="cell" draggable="true">{this.props.value}</td>)
  }
}

export default Cell;
