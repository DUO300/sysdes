/* placeholder file for JavaScript */
const confirm_delete = (id) => {
    if (window.confirm(`Task ${id} を削除します．よろしいですか？`)) {
        location.href = `/task/delete/${id}`;
    }
}

const confirm_update = (id) => {
    if (window.confirm(`Task ${id} を書き換えます．よろしいですか？`)) {
        document.getElementById(`update-form`).submit();
    }
}

const past_deadline_font_color = (date_str_all) => {
    var date_str = date_str_all.slice(0, 19);
    var date = new Date(Date.parse(date_str));
    var now = new Date(Date.now());
    if (date.getTime() < now.getTime()) {
        return `'#ff0000'`;
    }
    else {
        return `'#00ff00'`;
    }
}