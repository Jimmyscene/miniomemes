import React, { Component } from 'react';
import './App.css';
import Panel from './Panel.js';

class App extends Component {
  render() {
    return (
      <div class="parent">
        <div class="left">
          <Panel lang="golang"/>
        </div>
        <div class="right">
          <Panel lang="python"/>
        </div>
      </div>
     );
  }
}

export default App;
