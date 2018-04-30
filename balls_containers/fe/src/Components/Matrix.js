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
        this.props.onArrange();
    }

    render() {
        console.info("this.props.matrix", this.props.matrix);
        if (!this.props.matrix) {
            console.info("no matrix, props:", this.props);
            return ""
        }
        let rows = this.props.matrix.map((r, rowNumber) => {
            let matrix=this.props.index;
            let key = matrix.toString() + "_" + rowNumber.toString();
            return (
                <Row
                    cells={r}
                    row={rowNumber}
                    onDrop={this.props.onDrop.bind(this)}
                    matrix={this.props.index}
                        key={key}
                />
            )
        });
        console.info("rows", rows);
        let buttons = (
            <button className="btn btn-danger" onClick={this.arrange.bind(this)}>Arrange</button>
        )
        return (
            <div className="matrix">
                <div className="info" dangerouslySetInnerHTML={{__html: this.props.info}}></div>
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
