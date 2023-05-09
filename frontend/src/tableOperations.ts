import WatchList from "./interfaces/WatchList";

function showHideRow(wl: WatchList) {
    const row: HTMLElement | null = document.getElementById(`hidden-row-${wl.id}`)
    if (row!==null && wl.shortNote.length>0) {
        row.style.display = row.style.display==="none" ? "" : "none"
    }
}


export {showHideRow}