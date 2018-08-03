import React, { Component } from "react";

class Counter extends Component {
  state = {
    count: this.props.value,
    imageUrl: "https://laracasts.com/images/series/circles/do-you-react.png"
  };

  handleIncrement = () => {
    this.setState({ count: this.state.count + 1 });
  };

  render() {
    return (
      <div>
        {this.props.children}
        <img
          src={this.state.imageUrl}
          alt=""
          width="25"
          height="25"
          className="m-2"
        />
        <span style={this.styles} className={this.getBadgeClasses()}>
          {this.formatCount()}
        </span>
        <button
          onClick={this.handleIncrement}
          className="btn btn-secondary btn-sm m-2"
        >
          Increment
        </button>
      </div>
    );
  }

  getBadgeClasses() {
    let classes = "badge m-2 badge-";
    classes += this.state.count === 0 ? "warning" : "primary";
    return classes;
  }

  formatCount() {
    let { count } = this.state;
    return count === 0 ? "Zero" : count;
  }
}

export default Counter;
