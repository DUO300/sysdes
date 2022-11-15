/* placeholder file for JavaScript */
const confirm_delete = (id) => {
    if (window.confirm(`Task ${id} を削除します．よろしいですか？`)) {
        location.href = `/task/delete/${id}`;
    }
}

const confirm_delete_user = (name) => {
    if (window.confirm(`アカウント(アカウント名: ${name}) を削除します．よろしいですか？`)) {
        location.href = `/user/delete`;
    }
}

const confirm_update = (id) => {
    if (window.confirm(`Task ${id} を書き換えます．よろしいですか？`)) {
        document.getElementById(`update-form`).submit();
    }
}

const confirm_update_user = (name) => {
    if (window.confirm(`アカウント情報を更新します．よろしいですか？`)) {
        document.getElementById(`update-form`).submit();
    }
}

const past_deadline_font_color = (date_str_all, id) => {
    console.log(typeof date_str_all);
    console.log(date_str_all);
    var date_str = date_str_all.slice(0, 19);
    var date = new Date(Date.parse(date_str));
    var now = new Date(Date.now());
    if (date.getTime() > now.getTime()) {
        document.getElementById(id).style.color = "black";
    }
    else {
        document.getElementById(id).style.color = "red";
    }
}