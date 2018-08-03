import React, { Component } from "react";

class Data extends Component {
  state = {};
  render() {
    return (
      <div>
        User ID : {this.props.userId}
        <br />
        Title ID : {this.props.titleId}
        <br />
        Title : {this.props.title}
      </div>
    );
  }
}

export default Data;
