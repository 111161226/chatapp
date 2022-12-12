/* placeholder file for JavaScript */
//logout alert
const confirm_logout = () => {
    if(window.confirm(`ログアウトします．よろしいですか？`)) {
        location.href = `/logout`;
    }
}

//withdraw alert
const confirm_withdraw = () => {
    if(window.confirm(`このアプリを退会します. よろしいですか？`)) {
        location.href = `/withdraw`;
    }
}

//update user information alert 
const confirm_user_update = () => {
    if(window.confirm(`ユーザー情報を更新します．よろしいですか？`)) return true;
    else return false;
}
