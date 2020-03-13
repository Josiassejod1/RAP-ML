import React, { Component } from "react";
import axios from "axios";
import logo from "./logo.svg";
import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css";
import Button from "react-bootstrap/Button";
import Alert from "react-bootstrap/Alert";
import ToggleButton from "react-bootstrap/ToggleButton";
import ToggleButtonGroup from "react-bootstrap/ToggleButtonGroup";
import validator from 'validator';

class App extends Component {
  state = {
    image: "",
    query: "",
    submitBtn: false,
    uploadStatus: "",
    value: 1,
    toggleUploadAlert: false
  };

  handleUpdate(e) {
    console.log(e.target.value);
    this.setState({
      query: e.target.value
    });
  }

  handleUrl(e) {
    this.setState({
      image: e.target.value
    });
  }

  handleRadio(e) {
    this.setState({
      value: e.target.value
    });
  }

  getInfo() {
    const API_URL = `http://localhost:3002/artist`;

    axios
      .get(`${API_URL}`, {
        params: {
          artist: `${this.state.query}`
        }
      })
      .then(response => {
        this.setState({
          image: response.data,
          submitBtn: true,
        });
        console.log(response.data);
      })
      .catch(error => console.log(error));
  }

  validateURL(url) {
    var sanatizedUrl = url.toString()
    console.log(sanatizedUrl)
    if (validator.isURL(sanatizedUrl, {
      protocols: ['http', 'https'], 
      require_protocol:true,
      allow_underscores: true})) {
      this.setState({
        image: sanatizedUrl,
        submitBtn: true,
      })
    } else {
      console.log("NOPE")
      this.setState({
        toggleUploadAlert: true
      })
    }
  }

  uploadImage(url) {
    console.log(url);
    const API_URL = `http://localhost:3002/upload`;

    axios({
      url: `${API_URL}`,
      data: {
        image: `${this.state.image}`,
        metadata: `${this.state.query}`
      },
      method: "POST",
      headers: { "content-type": "application/x-www-form-urlencoded" }
    })
      .then(response => {
        this.setState({
          query: "",
          submitBtn: false,
          image: "",
          uploadStatus: response.data
        });
        console.log(response);
      })
      .catch(error => console.log(error));
  }

  render() {
    let img;
    let upload;
    let alert;
    let status = this.state.uploadStatus;
    let timer = null;
    let uploadAlert = this.state.toggleUploadAlert;

    let activeSearch;
    let alert2;

    if (this.state.value == 1) {
      activeSearch = (
        <div className="Center" id="artistSearch">
          <input
            autoFocus
            className="form-control"
            style={{ width: "500px", margin: "auto" }}
            type="text"
            size="lg"
            placeholder="Search for Artist 🎤"
            onChange={e => this.handleUpdate(e)}
            required
          />
        </div>
      );
    } else {
      activeSearch = (
        <div className="Center" id="urlSearch">
          <input
            autoFocus
            className="form-control"
            style={{ width: "500px", margin: "auto" }}
            type="text"
            size="lg"
            placeholder="Upload Image Url 🔗"
            onChange={e => this.handleUrl(e)}
            required
          />
        </div>
      );
    }
    if (uploadAlert) {
      alert2 = <Alert variant="warning">Image Url is Invalid: [http, https] Needed</Alert>
      setTimeout(() => {
        this.setState({
          toggleUploadAlert: false
        });
      }, 7000);
    }
    if (status != "") {
      if (status == "Image Successfully Uploaded") {
        alert = <Alert variant="success">{status}</Alert>;
      } else {
        alert = <Alert variant="warning">Opps ! Something went wrong. Try Pasting A URL Instead Using The Image Link Tab.</Alert>;
      }
      timer = setTimeout(() => {
        this.setState({
          uploadStatus: ""
        });
      }, 6000);
    } else {
      alert = <Alert></Alert>;
    }
    if (this.state.submitBtn) {
      img = (
        <img
          src={this.state.image || "unknown.jpg"}
          width="250"
          height="250"
          style={{
            borderRadius: "50%",
            padding: "25px"
          }}
        />
      );
      upload = (
        <Button
          onClick={e => this.state.value == 1 ? this.uploadImage(this.state.query) : this.uploadImage(this.state.image)}
          disabled={this.state.image == "unknown.jpg"}
        >
          Upload
        </Button>
      );
    } else {
      img = <img></img>;
      upload = <div></div>;
    }
    return (
      <div className="App">
        <header className="App-header">
          {alert2}
          {alert}
          {activeSearch}

          <ToggleButtonGroup type="radio" name="options" onClick={e => this.handleRadio(e)}>
            <ToggleButton variant='dark' value={1}>Keyword Search</ToggleButton>
            <ToggleButton variant='light' value={2}>Image Link</ToggleButton>
          </ToggleButtonGroup>
          
          {img}
          <br/>
          <div style={{ display: "inline", margin: "50px" }}>
            {upload}
            <Button variant="success" type="button" 
            onClick={e => this.state.value == 1 ? this.getInfo(e) : this.validateURL(this.state.image)}
            disabled={this.state.query === '' && this.state.image === ''}>
              Submit
            </Button>
          </div>
        </header>
      </div>
    );
  }
}

export default App;
