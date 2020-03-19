import React, { Component } from "react";
import axios from "axios";
import logo from "./logo.svg";
import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css";
import Button from "react-bootstrap/Button";
import Tab from "react-bootstrap/Tab";
import Alert from "react-bootstrap/Alert";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Nav from "react-bootstrap/Nav";
import ToggleButton from "react-bootstrap/ToggleButton";
import ToggleButtonGroup from "react-bootstrap/ToggleButtonGroup";
import validator from "validator";
import { FormControl } from "react-bootstrap";

class App extends Component {
  constructor() {
    super();
    this.input = React.createRef();
    this.boundBox = this.boundBox.bind(this);
    this.drawBox = this.drawBox.bind(this);
    this.producePrediction = this.producePrediction.bind(this);
  }
  state = {
    image: "",
    query: "",
    submitBtn: false,
    uploadStatus: "",
    value: 1,
    toggleUploadAlert: false,
    selectedFile: null,
    imageDetails: {
      width: '',
      height: '',
      clarifaiFaces: new Array(0),
      realFaces: new Array(0),
      blob: '',
      predictions: new Array(0)
    }
  };

  fileSelectedHandler = event => {
    var canvas = document.getElementById('canvas')
    if (canvas.style.backgroundImage !== null) {
      var ctx = canvas.getContext("2d")
      ctx.clearRect(0,0, canvas.width, canvas.height)
      canvas.style.backgroundImage = ""
    }

    var selectedFile = event.target.files[0]
    this.setState({
      selectedFile: selectedFile,
    })

    var reader = new FileReader();
    reader.readAsDataURL(selectedFile);
    if (selectedFile != null) {
      reader.onload = () => {
        this.state.imageDetails.blob = reader.result
        canvas.style.backgroundImage = `url(${reader.result})`
        canvas.style.backgroundSize = 'contain'
        canvas.style.backgroundRepeat = 'no-repeat'
      }
    }
  }

  fileUploaderHander = () => {
    const API_URL = `http://localhost:3002/predict`;
    var encoded =  this.state.selectedFile;
    var blob = this.state.imageDetails.blob
   if(encoded != null) {
      var base64 = blob.replace(/^data:image\/(.*);base64,/, '')
      axios
      .get(`${API_URL}`, {
        params: {
          image: base64
        },
        onUploadProgress: ProgressEvent => {
          console.log("Upload Progress: " + ((ProgressEvent.loaded / ProgressEvent.total) * 100))
        }
      }).then(response => {
        console.log(response);
        this.boundBox(blob, response)
      })
      .catch(error => console.log(error));
   }
  }

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
    this.searchField.value = "";
    this.setState({
      value: e.target.value,
      query: "",
      submitBtn: false
    });
  }

  drawBox(){
    var canvas = document.getElementById("canvas")
    var img = document.getElementById("predictionImage")
    img.style.display = 'none'
    var width = this.state.imageDetails.width
    var height =  this.state.imageDetails.height
    var imageDetails = this.state.imageDetails
    canvas.style.backgroundSize = 'contain'
    canvas.style.backgroundRepeat = 'no-repeat'
    canvas.width = width
    canvas.height = height
   


    var box;

    for(var i=0; i<imageDetails.clarifaiFaces.length; i++) {  
      console.log(imageDetails.clarifaiFaces[i].left_col)
      console.log(imageDetails.clarifaiFaces[i].top_row)
      console.log(imageDetails.clarifaiFaces[i].right_col)
      console.log(imageDetails.clarifaiFaces[i].bottom_row)
      box = {
        x: (imageDetails.clarifaiFaces[i].left_col * parseInt(canvas.width)),
        y: (imageDetails.clarifaiFaces[i].top_row * parseInt(canvas.height)),
        w: (imageDetails.clarifaiFaces[i].right_col * parseInt(canvas.width)) - (imageDetails.clarifaiFaces[i].left_col * parseInt(canvas.width)),
        h: (imageDetails.clarifaiFaces[i].bottom_row * parseInt(canvas.height)) - (imageDetails.clarifaiFaces[i].top_row * parseInt(canvas.height))
      }

  
        var ctx = canvas.getContext("2d")
       // ctx.style.backgroundImage = 
        // ctx.width = width
        // ctx.height = height
        //ctx.drawImage(img, 0,0, width, height)
       
  
        ctx.textBaseline = "top"
        imageDetails.realFaces.push(box)
        ctx.font = (box.w * 1.4) + "px monospace";
        ctx.fillText("🎤", box.x - (box.w / 5), box.y - (box.h/4));
      
   
    
    }

   

    


  }

  producePrediction(data){
    console.log(data)
    var artist = new Array(0);
    if (data != null) {
      var prediction =  data
      console.log(prediction)
      for (var i = 0; i < prediction.length; i++) {
        var concepts = prediction[i].data.concepts
        console.log(concepts)
        for(var j = 0; j < concepts.length; j++) {
          console.log(concepts[i][j])
          artist.push({
            name: concepts[j].name,
            value: concepts[j].value
          })
        }
      }
      console.log(artist)
      this.state.imageDetails.predictions = artist
    } else {
      console.log("Data empty 🙁")
    }
  }


  
  
  boundBox(base64, res){
    var prediction =  document.getElementById("predictionImage")
    prediction.src = base64
   
   

    this.state.imageDetails.width =  prediction.width 
    this.state.imageDetails.height=  prediction.height

    var data = res.data.outputs[0].data.regions;
    this.producePrediction(data)
    console.log(data)
    var regionBox = new Array(0);

    if (data !== null) {
      for (var i = 0; i < data.length; i++) {
        regionBox.push(data[i].region_info.bounding_box)
      }
      this.state.imageDetails.clarifaiFaces = regionBox 

      this.drawBox()

    } else {
      console.log("error")
    }

   
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
          submitBtn: true
        });
        console.log(response.data);
      })
      .catch(error => console.log(error));
  }

 
  validateURL(url) {
    var sanatizedUrl = url.toString();
    console.log(sanatizedUrl);
    if (
      validator.isURL(sanatizedUrl, {
        protocols: ["http", "https"],
        require_protocol: true,
        allow_underscores: true
      })
    ) {
      this.setState({
        image: sanatizedUrl,
        submitBtn: true
      });
    } else {
      console.log("NOPE");
      this.setState({
        toggleUploadAlert: true
      });
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
    this.searchField.value = "";
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

    activeSearch = (
      <div className="Center">
        <input
          ref={el => (this.searchField = el)}
          autoFocus
          className="form-control"
          style={{ width: "500px", margin: "auto" }}
          type="text"
          size="lg"
          placeholder={
            this.state.value == 1 ? "Search for Artist 🎤" : "Add Image Url 🔗"
          }
          onChange={e =>
            this.state.value == 1 ? this.handleUpdate(e) : this.handleUrl(e)
          }
          required
        />
      </div>
    );

    if (uploadAlert) {
      alert2 = (
        <Alert variant="warning">
          Image Url is Invalid: [http, https] Needed
        </Alert>
      );
      setTimeout(() => {
        this.setState({
          toggleUploadAlert: false
        });
      }, 10000);
    }
    if (status != "") {
      if (status == "Image Successfully Uploaded") {
        alert = <Alert variant="success">{status}</Alert>;
      } else {
        alert = <Alert variant="warning">{status}</Alert>;
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
          onClick={e =>
            this.state.value == 1
              ? this.uploadImage(this.state.query)
              : this.uploadImage(this.state.image)
          }
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
      <Tab.Container id="left-tabs-example" defaultActiveKey="first">
        <Row>
          <Col sm={3}>
            <Nav  variant="pills" className="flex-column">
              <Nav.Item>
                <Nav.Link eventKey="first">Train</Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link eventKey="second">Prediction</Nav.Link>
              </Nav.Item>
            </Nav>
          </Col>

          <Col sm={9}>
            <Tab.Content>
              <Tab.Pane eventKey="first">
                <div className="App">
                  <header className="App-header">
                    {alert2}
                    {alert}
                    {activeSearch}

                    <ToggleButtonGroup
                      type="radio"
                      name="options"
                      onClick={e => this.handleRadio(e)}
                    >
                      <ToggleButton variant="dark" value={1}>
                        Keyword Search
                      </ToggleButton>
                      <ToggleButton variant="light" value={2}>
                        Image Link
                      </ToggleButton>
                    </ToggleButtonGroup>

                    {img}
                    <br />
                    <div style={{ display: "inline", margin: "50px" }}>
                      {upload}
                      <Button
                        variant="success"
                        type="button"
                        onClick={e =>
                          this.state.value == 1
                            ? this.getInfo(e)
                            : this.validateURL(this.state.image)
                        }
                        disabled={
                          this.state.query === "" && this.state.image === ""
                        }
                      >
                        Submit
                      </Button>
                    </div>
                  </header>
                </div>
              </Tab.Pane>
              <Tab.Pane eventKey="second">
              <div className="App">
                  <header className="App-header">
                    <div>
                    <img id="predictionImage" src=""></img>
                    <canvas id="canvas"></canvas>
                    </div>
                    <div style={{backgroundColor: "gray"}}>
                    <input  
                    type="file" 
                    style={{display: 'none'}}
                    onChange={this.fileSelectedHandler}
                    ref={fileInput => this.fileInput = fileInput }
                    />
                      <Button onClick={() => this.fileInput.click()} variant="success">Pick File</Button>
                      <Button onClick={this.fileUploaderHander} variant="success">Upload</Button>
                    </div>
                  </header>
                </div>
              </Tab.Pane>
            </Tab.Content>
          </Col>
        </Row>
      </Tab.Container>
    );
  }
}

export default App;
