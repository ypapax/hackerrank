import React, {Component} from 'react';

class Cell extends Component {
    allowDrop(e) {
        e.preventDefault();
    }

    onDrop() {
        let column = this.props.column;
        debugger;
        this.props.onDrop(column);
    }

    render() {
        return (
            <td className="cell" draggable="true" onDrop={this.onDrop.bind(this)}
                onDragOver={this.allowDrop.bind(this)}>
                {this.props.value}
            </td>
        )
    }
}

export default Cell;
