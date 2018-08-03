import React, { Component } from "react";

class Counter extends Component {
  state = {
    imageUrl: "https://laracasts.com/images/series/circles/do-you-react.png"
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
          onClick={() => this.props.onIncrement(this.props.counter)}
          className="btn btn-secondary btn-sm m-2"
        >
          Increment
        </button>
        <button
          onClick={() => this.props.onDelete(this.props.counter.id)}
          className="btn btn-danger btn-sm m-2"
        >
          Delete
        </button>
      </div>
    );
  }

  getBadgeClasses() {
    let classes = "badge m-2 badge-";
    classes += this.props.counter.value === 0 ? "warning" : "primary";
    return classes;
  }

  formatCount() {
    let { value } = this.props.counter;
    return value === 0 ? "Zero" : value;
  }
}

export default Counter;
