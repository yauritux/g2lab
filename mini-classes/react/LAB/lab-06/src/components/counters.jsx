import React, { Component } from "react";

import Counter from "./counter";

class Counters extends Component {
  state = {
    counters: [{ id: 1, value: 5 }, { id: 2, value: 0 }, { id: 3, value: 0 }]
  };

  handleDelete = counterId => {
    const counters = this.state.counters.filter(c => c.id !== counterId);
    this.setState({ counters });
  };

  render() {
    return (
      <div>
        {this.state.counters.map(counter => (
          <Counter
            key={counter.id}
            counter={counter}
            onDelete={this.handleDelete}
          >
            <h4>Counter #{counter.id}</h4>
          </Counter>
        ))}
      </div>
    );
  }
}

export default Counters;
