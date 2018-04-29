import React, { Component } from 'react';
import './App.css';
import Arranger from "./Components/Arranger";
import $ from "jquery";

const backend = "http://localhost:8083";

class App extends Component {

  onArrange(input){
      $.ajax({
          type: "POST",
          url: backend + "/api/v1/arrange",
          data: JSON.stringify({Params:[input]}),
          success: function(data) {
              console.info("data", data);
          },
          dataType: "json",
          error: function(e) {
              console.error("Could not request expenses data");
          }
      });
  }
  render() {
    return (
      <div className="App">
          <Arranger
            onArrange={this.onArrange}/>
      </div>
    );
  }
}

export default App;
