import React, { Component } from "react";

import Data from "./data";

import axios from "axios";

class Table extends Component {
  state = {
    records: {}
  };

  constructor() {
    super();
    console.log("constructor");
    (async () => {
      let data = await fetchData();
      console.log(data);
      this.setState({
        records: { user_id: data.userId, title_id: data.id, title: data.title }
      });
    })();
  }
  render() {
    return (
      <React.Fragment>
        <ul>
          <Data
            userId={this.state.records.user_id}
            titleId={this.state.records.title_id}
            title={this.state.records.title}
          />
        </ul>
      </React.Fragment>
    );
  }
}

function fetchData() {
  return new Promise((resolve, reject) => {
    axios
      .get("https://jsonplaceholder.typicode.com/posts/1")
      .then(rows => {
        resolve(rows.data);
      })
      .catch(err => {
        reject(err.message);
      });
  });
}

export default Table;
