import React, { Component } from 'react';
import axios from 'axios';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  state = {
    image: '',
    query: ''
  }

  handleUpdate(e) {
    console.log(e.target.value)
    this.setState({
      query: e.target.value
    })
  } 


getInfo() {
  

  const API_URL = `http://localhost:3002/artist`

  axios.get(`${API_URL}`, {
    params: {
      artist:  `${this.state.query}`
    }
  })
  .then(response => {
    this.setState({ 
      image: response.data
    })
    console.log(response.data)
  })
  .catch(error => console.log(error));
}

uploadImage(url) {
  
  console.log(url)
  const API_URL = `http://localhost:3002/upload`

  axios({
    url: `${API_URL}`,
     data: {
      image:  `${this.state.image}`,
      metadata: `${this.state.query}`
     },
     method: 'POST',
     headers: { 'content-type': 'application/x-www-form-urlencoded' }

  })
  .then(response => {
    console.log(response)
  })
  .catch(error => console.log(error));
}

  
  render(){
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.js</code> and save to reload.
          </p>
          <div className="Center">
          <input
            className="form-control"
            style={{width: '300px', margin: 'auto'}}
            type="text"
            size="lg"
            placeholder="Search for Lyrics 🔎"
            onChange={(e) => this.handleUpdate(e)}
            required
          />
          </div>
          <input type="button" value="Submit" onClick={(e) => this.getInfo(e)}/>
          <img src={this.state.image || 'unknown.jpg'} width="250" height="250" />

          <button onClick={(e) => this.uploadImage(this.state.query)}>Upload</button>
        </header>
      </div>
    );
  }
}
 


export default App;
