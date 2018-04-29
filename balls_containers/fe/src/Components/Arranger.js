import React, { Component } from 'react';
import $ from 'jquery'
class Arranger extends Component {
    constructor(){
        super();
        this.state = {
            input: `1
4
997612619 934920795 998879231 999926463
960369681 997828120 999792735 979622676
999013654 998634077 997988323 958769423
997409523 999301350 940952923 993020546`
        }
    }
    componentDidMount(){

    }
  render() {
    return (
      <div className="arranger">
          <div>
              <textarea rows="10" cols="10" value={this.state.input}>
                </textarea>
          </div>

          <button className="btn btn-success">Arrange</button>
      </div>
    );
  }
}

export default Arranger;
