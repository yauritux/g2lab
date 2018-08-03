import React, { Component } from "react";

import Data from "./data";

import axios from "axios";
//import _ from "lodash";

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
    //console.log("records in render:");
    //console.log(this.state.records);
    //{this.state.records.map(r => <Data key={r.id} userId={r.user_id} />)}
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
        //console.log("rows:");
        //console.log(rows.data);
        resolve(rows.data);
      })
      .catch(err => {
        //console.error("error:", err.message);
        reject(err.message);
      });
  });
}

export default Table;
