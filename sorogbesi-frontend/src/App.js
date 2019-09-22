import React, { Component} from 'react';
import './App.css';
import {connectViaWebSocket, sendMessageViaWebSocket} from './api'
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';
import ChatInput from './components/ChatInput/ChatInput';

class App extends Component{
  // eslint-disable-next-line 
  constructor(props){
    super(props);

    // define the initial state for chatHistory
    this.state = {
      chatHistory: []
    }
  }

  //this.connectViaWebSocket = this.connectViaWebSocket.bind(this);
  //this.send = this.send.bind(this);

  send(event){
    console.log("Client: before sending message via websocket, to the Server");
    if (event.keyCode === 13){
      sendMessageViaWebSocket(event.target.value);
      event.target.value = "";
    }
  }

  // using the message eched from the Websocket server, to update the application state in the frontend
  componentDidMount() {
    connectViaWebSocket((msg) => {
      console.log("New message is being added to chat history ... ")
      this.setState(prevState => ({
//        chatHistory: [...this.state.chatHistory, msg] // new msg is appended to the currentState of the chatHistory
        chatHistory: [...prevState.chatHistory, msg] // new msg is appended to the currentState of the chatHistory
      }))
      console.log(this.state);
    });
  }

  render(){
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
      </div>
    )
  }
}

export default App;
