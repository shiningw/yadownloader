import helper from '../utils/helper'
import eventHandler from '../lib/eventHandler'
import Clipboard from '../lib/clipboard'
import '../css/clipboard.scss';

const buttonHandler = (event, type) => {
    let element = event.target;
    event.stopPropagation();
    event.preventDefault();
    let url = helper.generateUrl(element.getAttribute("path"));
    let row, data = {};
    let removeRow = true;
    if (element.getAttribute("id") == "download-action-button") {
        helper.loop(helper.getCounters);
        helper.setContentTableType("search-results");
    }
    if (row = element.closest('.table-row-search')) {
        if (element.className == 'icon-clipboard') {
            const clippy = new Clipboard(element, row.dataset.link);
            clippy.Copy();
            return;
        }
        data['text-input-value'] = row.dataset.link;
    } else {
        row = element.closest('.table-row')
        data = row.dataset;
        if (!data.gid) {
            console.log("gid is not set!");
        }
    }
    helper.httpClient(url).setErrorHandler(function (xhr, textStatus, error) {
        console.log(error);
    }).setHandler(function (data) {
        if (data.hasOwnProperty('error') && data.error) {
            helper.error(data['error']);
            return;
        }
        if (data.hasOwnProperty('result') && data.result) {
            helper.message("Success " + data['result']);
        }
        if (data.hasOwnProperty('data') && data.data) {
            helper.message(data.data);
        }
        if (row && removeRow)
            row.remove();
    }).setData(data).send();

}
export default {
    run: function () {
        eventHandler.add("click", "#downloader-table-wrapper", ".table-cell-action-item .button-container button", e => buttonHandler(e, ''));
        eventHandler.add("click", "#downloader-table-wrapper", ".table-row button.icon-clipboard", function (e) {
            let element = e.target;
            const clippy = new Clipboard(element);
            clippy.Copy();
        });
    }
}