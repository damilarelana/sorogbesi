import React, { Component } from "react";
import "./Message.scss";

class Message extends Component{
  constructor(props){
    super(props);

    // received the JSON message and store it in the state variable message
    let temp = JSON.parse(this.props.message);
    this.state = {
      message: temp
    };
  }

  render() {
    return(
      <div className="Message">
        {this.state.message.body}
      </div>
    );
  }
}

export default Message;