import React, { Component } from 'react';
import logo from './logo.svg';

class Panel extends Component {
    constructor(props) {
        super(props);
        this.state = {
            data: {}
        }
    }
    componentDidMount() {
            fetch(this.props.lang)
              .then(response => response.json())
              .then(data => this.setState({ data }))
              .catch(error => {
                  console.log(error)
              })

    }
    render() {
        return (
            <div className="App">
                <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" />
                <p>This is the {this.props.lang} panel</p>
                {
                    Object.entries(this.state.data).map(([bucket, items]) => {
                        return (
                            <div>
                                <p>Bucket: {bucket} </p>
                                {
                                    Object.entries(items).map(([name, url]) => {
                                        return (
                                            <span>Item: <a style={{color: 'white'}} href={url}>{name}</a><br/></span>
                                        )
                                    })
                                }
                            </div>
                        )
                    })
                }
                </header>
            </div>
        )
    }
}
export default Panel;
