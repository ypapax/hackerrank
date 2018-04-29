import React, { Component } from 'react';
import Matrix from "./Matrix";
class Arranger extends Component {
    constructor(props){
        super(props);
        this.state = {
            input: `1
4
997612619 934920795 998879231 999926463
960369681 997828120 999792735 979622676
999013654 998634077 997988323 958769423
997409523 999301350 940952923 993020546`,
            matrix: props.matrix
        };
    }
    componentDidMount(){

    }
    onParse(){
        this.props.onParse(this.state.input);
    }
    inputChanged(input){
        this.setState({input: input});
    }
  render() {
        console.info("Arranger render, arranger state", this.state);
    return (
      <div className="arranger">
          <div>
              <textarea rows="10" cols="10" value={this.state.input}  onChange={this.inputChanged.bind(this)}>
              </textarea>
          </div>

          <button className="btn btn-success" onClick={this.onParse.bind(this)}>Parse</button>
          <Matrix matrix={this.props.matrix}/>
      </div>
    );
  }
}

export default Arranger;
