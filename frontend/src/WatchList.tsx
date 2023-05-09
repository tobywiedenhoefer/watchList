import * as React from "react";
import {useEffect, useState} from "react";

import api from "./api";
import WatchList from "./interfaces/WatchList";
import { showHideRow } from "./tableOperations";
import { StringOrNumber } from "./react-app-env";

import "./style/WatchList.css"
import WatchListRowShownComponent from "./WatchListTableRowShow";
import WatchListRowHiddenComponent from "./WatchListRowHiddenComponent";

function WatchListComponent (): JSX.Element {
    const [watchList, setWatchList] = useState<WatchList[]>([])
    useEffect(() => {
        const getWatchList = async () => {
            let jsonBody: WatchList[]
            try {
                jsonBody = await api.get<WatchList[]>()
            } catch (e) {
                console.log(e)
                jsonBody = []
            }
            setWatchList(jsonBody)
        }
        if (watchList.length===0) {
            getWatchList()
        }
    }, [])
    return (
        <>
            <table id="watchListTable" className="WatchList">
                <thead>
                <tr>
                    <th className="title">Title</th>
                    <th>Media Type</th>
                    <th>Genre</th>
                    <th>Steaming Platform</th>
                </tr>
                </thead>
                <tbody>
                {watchList.map((watchListItem) => {
                    return (
                        <>
                            <WatchListRowShownComponent {...watchListItem} />
                            <WatchListRowHiddenComponent {...watchListItem} />
                        </>
                    )
                })}
                </tbody>
                <tfoot>

                </tfoot>
            </table>
        </>
    )
}

export default WatchListComponent
