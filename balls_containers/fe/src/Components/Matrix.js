import React, {Component} from 'react';
import Row from "./Row";

class Matrix extends Component {
    constructor(props) {
        super(props);
        this.state = {
            "matrix": props.matrix
        };
        console.info("props.matrix", props.matrix);
    }

    arrange() {
        this.props.arrange();
    }

    onDrop(row, column){
        this.props.onDrop(row, column);
    }

    render() {
        console.info("this.props.matrix", this.props.matrix);
        if (!this.props.matrix) {
            return ""
        }
        let rows = this.props.matrix.map((r, index) => {
            return (
                <Row
                    cells={r}
                    row={index}
                    onDrop={this.onDrop.bind(this)}
                />
            )
        });
        console.info("rows", rows);
        let buttons = (
            <button className="btn btn-danger" onClick={this.arrange.bind(this)}>Arrange</button>
        )
        return (
            <div className="matrix">
                <table className="table">
                    <tbody>
                        {rows}
                    </tbody>
                </table>
                {buttons}
            </div>
        );
    }
}

export default Matrix;
