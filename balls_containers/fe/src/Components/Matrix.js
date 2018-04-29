import React, { Component } from 'react';
import Row from "./Row";

class Matrix extends Component {
    constructor(props){
        super(props);
        this.state = {
            "matrix": props.matrix
        };
        console.info("props.matrix", props.matrix);
    }

  render() {
      console.info("this.props.matrix", this.props.matrix);
        if (!this.props.matrix) {
            return ""
        }
        let rows = this.props.matrix.map((r, index) => {
            return (<Row cells={r} key={index}/>)
        });
        console.info("rows", rows);
    return (
      <div className="matrix">
          {rows}
      </div>
    );
  }
}

export default Matrix;
