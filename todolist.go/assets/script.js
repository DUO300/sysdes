const confirm_delete_task = (id) => {
    if (window.confirm(`Task ${id} を削除します．よろしいですか？`)) {
        location.href = `/task/delete/${id}`;
    }
}

const confirm_delete_user = (name) => {
    if (window.confirm(`アカウント(ユーザ名: ${name}) を削除します．よろしいですか？`)) {
        location.href = `/user/delete`;
    }
}

const confirm_update_task = (id) => {
    if (window.confirm(`Task ${id} を書き換えます．よろしいですか？`)) {
        document.getElementById(`update-form`).submit();
    }
}

const confirm_update_user = () => {
    if (window.confirm(`ユーザ情報を更新します．よろしいですか？`)) {
        document.getElementById(`update-form`).submit();
    }
}

const parse_go_formatted_date = (date_str_go_formatted) => {
    if (date_str_go_formatted.length < 19) {
        console.error("Unknown date format");
        return -1;
    }
    const date_str = date_str_go_formatted.substring(0, 10) + "T" + date_str_go_formatted.substring(11, 19);
    return new Date(Date.parse(date_str));
}

const past_deadline_bgcolor = (date_str_all, id, status) => {
    const date = parse_go_formatted_date(date_str_all);
    const now = new Date(Date.now());
    if (date.getTime() < now.getTime()) {
        if (status == "false") {
            // Set background color to red
            document.getElementById(id).style.backgroundColor = "#ffcccc";
        }
        else {
            // Set background color to green
            document.getElementById(id).style.backgroundColor = "#99ffcc";
        }
    }
}