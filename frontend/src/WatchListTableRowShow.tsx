import React from "react";
import WatchList from "./interfaces/WatchList";
import {showHideRow} from "./tableOperations";

const WatchListRowShownComponent = (watchListItem: WatchList): JSX.Element => {
    return (
        <tr
            key={watchListItem.id}
            id={`main-row-${watchListItem.id}`}
            onClick={() => showHideRow(watchListItem)}
            style={{cursor: `${watchListItem.shortNote.length>0 ? "pointer" : "default"}`}}
        >
            <td className="title">{watchListItem.title}</td>
            <td>{watchListItem.mediaType}</td>
            <td>{watchListItem.genre}</td>
            <td>{watchListItem.streamingPlatform}</td>
        </tr>
    )
}

export default WatchListRowShownComponent
