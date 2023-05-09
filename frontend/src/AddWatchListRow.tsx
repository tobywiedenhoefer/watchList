import React, {Dispatch, FormEvent, SetStateAction, useState} from "react";

import api from "./api";
import WatchList from "./interfaces/WatchList";
import WatchListRowShownComponent from "./WatchListTableRowShow";
import WatchListRowHiddenComponent from "./WatchListRowHiddenComponent";
import {getFormSubmissionInfo} from "react-router-dom/dist/dom";

const submitForm = (e:React.FormEvent<HTMLInputElement>): void => {
    e.preventDefault()
    let body: WatchList = {
        id: 0,
        title: e.currentTarget.title,
    }
    api.post<WatchList>()
}

const AddWatchListRow = (): JSX.Element => {
    const [watchList, setWatchList] = useState<WatchList[]>([])
    return (
        <>
            {watchList.map((watchListItem) => {
                return (
                    <>
                        <WatchListRowShownComponent {...watchListItem} />
                        <WatchListRowHiddenComponent {...watchListItem} />
                    </>
                )
            })}
            <tr id="add-table-row">
                <td><form id="add-row-form"><input id="formTitle" name="title" type="text" /></form></td>
                <td><input form="add-row-form" id="formMediaType" name="mediaType" type="number" /></td>
                <td><input form="add-row-form" id="formGenre"     name="genre" type="text" /></td>
                <td><input form="add-row-form" id="formStreamingPlatform" name="streamingPlatform" type="text" /></td>
            </tr>
            <tr id="submit-button-row">
                <td colSpan={3}></td>
                <td className="submit-row-addition" id="submit-row-addition">
                    <input form="add-row-form" type="button" onSubmit={submitForm} />
                </td>
            </tr>
        </>
    )
}

export default AddWatchListRow
