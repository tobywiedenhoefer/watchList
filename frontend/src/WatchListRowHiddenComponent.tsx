import React from "react";

import WatchList from "./interfaces/WatchList";

const WatchListRowHiddenComponent = (watchListItem: WatchList): JSX.Element => {
    return (
        <tr
            key={`hidden-row-${watchListItem.id}`}
            id={`hidden-row-${watchListItem.id}`}
            style={{display: "none"}}
            className="hidden-row"
        >
            <td colSpan={4} className="title hidden-row">
                {watchListItem.shortNote}
            </td>
        </tr>
    )
}

export default WatchListRowHiddenComponent
