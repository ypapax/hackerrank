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

    arrange(){
        this.props.arrange();
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
        let buttons = (
            <button className="btn btn-danger" onClick={this.arrange.bind(this)}>Arrange</button>
        )
    return (
      <div className="matrix">
          {rows}
          {buttons}
      </div>
    );
  }
}

export default Matrix;
