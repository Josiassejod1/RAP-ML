import React, { Component } from 'react';
import axios from 'axios';
import logo from './logo.svg';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Button from 'react-bootstrap/Button';
import Alert from 'react-bootstrap/Alert';
import ToggleButton from 'react-bootstrap/ToggleButton';
import ToggleButtonGroup from 'react-bootstrap/ToggleButtonGroup';

class App extends Component {
  state = {
    image: '',
    query: '',
    submitBtn: false,
    uploadStatus: ''
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
      image: response.data,
      submitBtn: true
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
    this.setState({
      query: '',
      submitBtn: false,
      image: '',
      uploadStatus: response.data
    })
    console.log(response)
  })
  .catch(error => console.log(error));
}

  
  render(){
    let img;
    let upload;
    let alert;
    let status = this.state.uploadStatus;
    let timer = null;
    if (status != ''){
      if (status == "Image Successfully Uploaded") {
      alert = <Alert variant="success">{status}</Alert>
      }  else {
        alert = <Alert variant="warning">Opps ! Something went wrong .</Alert>
      }
      timer = setTimeout(() => {
        this.setState({
          uploadStatus: ''
        })
      }, 3000)
    } else {
      alert = <Alert></Alert>
    }
    if (this.state.submitBtn) {
      img = <img src={this.state.image || 'unknown.jpg'} 
      width="250" 
      height="250" 
      style={{
        borderRadius: '50%',
        padding: '25px'
        }} />
      upload =  <Button onClick={(e) => this.uploadImage(this.state.query)} disabled={this.state.image == "unknown.jpg"}>Upload</Button>
    } else {
      img = <img></img>
      upload = <div></div>
    }
    return (
      <div className="App">
        <header className="App-header">
          {alert}
          <div className="Center">
          <input autoFocus
            className="form-control"
            style={{width: '500px', margin: 'auto'}}
            type="text"
            size="lg"
            placeholder="Search for Artist 🎤"
            onChange={(e) => this.handleUpdate(e)}
            required
          />
          </div>
          
      
          {img}
         
          <div style={{display: "inline"}}>
          {upload}
          <Button variant="dark" type="button" onClick={(e) => this.getInfo(e)}>Submit</Button>
          </div>
        
        
        </header>
      </div>
    );
  }
}
 


export default App;
