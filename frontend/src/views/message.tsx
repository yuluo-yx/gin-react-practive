import "./message.css";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import TextField from "@material-ui/core/TextField";
import InputAdornment from "@material-ui/core/InputAdornment";
import Button from "@material-ui/core/Button";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import ListItemAvatar from "@material-ui/core/ListItemAvatar";
import Typography from "@material-ui/core/Typography";
import getStr from "../utils/strings";
import React from "react";
import Avatar from "@material-ui/core/Avatar";
import { deepPurple } from "@material-ui/core/colors";
import IMsg from "../mock/messageData";
import { MsgPages, publishMsg, LoginApi } from "../api/api";
import { useEffect, useState } from "react";
import { getCurUserFromBrowser } from "../utils/user";
import Pagination from "@material-ui/lab/Pagination";
import Modal from "@material-ui/core/Modal";
import Backdrop from "@material-ui/core/Backdrop";
import Fade from "@material-ui/core/Fade";
import InputLabel from "@material-ui/core/InputLabel";
import FormControl from "@material-ui/core/FormControl";
import Visibility from "@material-ui/icons/Visibility";
import VisibilityOff from "@material-ui/icons/VisibilityOff";
import IconButton from "@material-ui/core/IconButton";
import Input from "@material-ui/core/Input";

interface State {
  username: string;
  password: string;
  showPassword: boolean;
}

interface User {
  userId: number;
  userName: string;
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    username: {
      width: "33vh",
    },
    modal: {
      display: "flex",
      alignItems: "center",
      justifyContent: "center",
    },
    paper: {
      backgroundColor: theme.palette.background.paper,
      border: "2px solid #000",
      boxShadow: theme.shadows[5],
      padding: theme.spacing(2, 4, 3),
    },
    root: {
      "& > *": {
        margin: theme.spacing(1),
        width: "70ch",
      },
    },
    iconButton: {
      padding: 10,
    },
    btn: {
      height: "6ch",
    },
    purple: {
      color: theme.palette.getContrastText(deepPurple[500]),
      backgroundColor: deepPurple[500],
      width: "5vh",
      height: "5vh",
      marginRight: "1vh",
    },
    inline: {
      display: "inline",
    },
  })
);

// const msgs: IMsg[] = [
//   {
//     id: 0,
//     user_id: 1,
//     user_name: "yuluo",
//     text: "祝你平安-yuluo！",
//   },
//   {
//     id: 1,
//     user_id: 2,
//     user_name: "huakai",
//     text: "祝你平安-yuluo！",
//   },
//   {
//     id: 2,
//     user_id: 3,
//     user_name: "Test",
//     text: "祝你平安-yuluo！",
//   },
// ];

export default function Message() {
  const [text, _] = useState("祝你平安！");

  const [values, setValues] = React.useState<State>({
    username: "",
    password: "",
    showPassword: false,
  });

  const handleChange =
    (prop: keyof State) => (event: React.ChangeEvent<HTMLInputElement>) => {
      setValues({ ...values, [prop]: event.target.value });
    };

  const handleClickShowPassword = () => {
    setValues({ ...values, showPassword: !values.showPassword });
  };

  const handleMouseDownPassword = (
    event: React.MouseEvent<HTMLButtonElement>
  ) => {
    event.preventDefault();
  };

  const [open, setOpen] = React.useState(false);

  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const [msgsDataTotal, setMsgDataTotal] = useState(0);
  const classes = useStyles();
  const [msgs, setMsgs] = useState<IMsg[]>([]);
  const GetMsgListAction = async (pageNum: number, pageSize: number) => {
    let res = await MsgPages({
      page_num: pageNum,
      page_size: pageSize,
    });
    // console.log("list", res);

    let msgR = res.data.item as IMsg[];
    setMsgs(msgR);
    setMsgDataTotal(res.data.total);
  };

  useEffect(() => {
    GetMsgListAction(1, 5);
    ShowUserAvatar();
  }, []);

  const listItem = msgs.map((item, _) => {
    return (
      <ListItem key={item.id} alignItems="flex-start">
        <ListItemAvatar>
          <Avatar className={classes.purple}>
            {getStr(0, 1, String(item.user_name))}
          </Avatar>
        </ListItemAvatar>
        <ListItemText
          primary={item.user_name}
          secondary={
            <React.Fragment>
              <Typography
                component="span"
                variant="body2"
                className={classes.inline}
                color="textPrimary"
              >
                {item.text}
              </Typography>
            </React.Fragment>
          }
        />
      </ListItem>
    );
  });

  const ShowUserAvatar = () => {
    let buser = JSON.parse(String(getCurUserFromBrowser()));
    // console.log("curUser browser: ", buser);
    if (buser === null) {
      // handleOpen();
    } else {
      curUser.userName = buser.user_name;
      setCurUser(curUser);
    }
  };

  const [currentPage, setCurrentPage] = useState(1);
  const messagePerPage = 5;
  const totalPages = Math.ceil(msgsDataTotal / messagePerPage);
  // console.log(totalPages);

  const handlePageChange = (_: any, value: number) => {
    setCurrentPage(value);
    GetMsgListAction(value, messagePerPage);
  };
  // const startIndex = (currentPage - 1) * messagePerPage;
  // const endIndex = startIndex + messagePerPage;

  const [curUser, setCurUser] = React.useState<User>({
    userId: -1,
    userName: "",
  });

  const getCurUser = () => {
    let buser = JSON.parse(String(getCurUserFromBrowser()));
    // console.log("curUser browser: ", buser);
    if (buser === null) {
      handleOpen();
    } else {
      curUser.userName = buser.user_name;
      curUser.userId = buser.id;
      setCurUser(curUser);
    }
  };

  const PublishMsg = async () => {
    // 获取登陆用户
    getCurUser();

    // 获取输入信息

    let res = await publishMsg({
      text: text,
      user_id: curUser.userId,
    });

    // console.log(res);

    if (res.status === 200) {
      alert("发布成功！");
      // 自动跳转到最新的一条数据
      window.location.reload();
    } else {
      alert("发布失败，请稍后再试....");
    }
  };

  const UserLogin = async () => {
    // console.log(values.username);
    // console.log(values.password);
    let res = await LoginApi({
      user_name: values.username,
      password: values.password,
    });

    // console.log(res);

    localStorage.setItem("user", JSON.stringify(res.data.user));
    localStorage.setItem("token", JSON.stringify(res.data.token));

    ShowUserAvatar();

    // data clear
    values.username = "";
    values.password = "";

    handleClose();
  };

  return (
    <div>
      <h3>Message Board</h3>
      <div id="msg">{listItem}</div>
      <br />
      <div id="pagination">
        <Pagination
          count={totalPages}
          page={currentPage}
          onChange={handlePageChange}
          variant="outlined"
          shape="rounded"
        />
      </div>
      <div id="input">
        <TextField
          className={classes.root}
          id="outlined-basic"
          label="Message"
          variant="outlined"
          defaultValue="祝你平安！"
          value={text}
          InputProps={{
            startAdornment: (
              <InputAdornment position="start">
                <Avatar className={classes.purple}>
                  {curUser.userName === ""
                    ? null
                    : getStr(0, 1, curUser.userName)}
                </Avatar>
              </InputAdornment>
            ),
          }}
        />
        <Button
          onClick={PublishMsg}
          className={classes.btn}
          variant="contained"
          color="primary"
        >
          发送
        </Button>
      </div>
      <Modal
        aria-labelledby="transition-modal-title"
        aria-describedby="transition-modal-description"
        className={classes.modal}
        open={open}
        onClose={handleClose}
        closeAfterTransition
        BackdropComponent={Backdrop}
        BackdropProps={{
          timeout: 500,
        }}
      >
        <Fade in={open}>
          <div className={classes.paper}>
            <h2 id="transition-modal-title">请登陆....</h2>
            <FormControl className={classes.username}>
              <InputLabel htmlFor="standard-adornment-username">
                Username
              </InputLabel>
              <Input
                id="standard-adornment-username"
                value={values.username}
                onChange={handleChange("username")}
              />
            </FormControl>
            <br />
            <FormControl>
              <InputLabel htmlFor="standard-adornment-password">
                Password
              </InputLabel>
              <Input
                id="standard-adornment-password"
                type={values.showPassword ? "text" : "password"}
                value={values.password}
                onChange={handleChange("password")}
                endAdornment={
                  <InputAdornment position="end">
                    <IconButton
                      aria-label="toggle password visibility"
                      onClick={handleClickShowPassword}
                      onMouseDown={handleMouseDownPassword}
                    >
                      {values.showPassword ? <Visibility /> : <VisibilityOff />}
                    </IconButton>
                  </InputAdornment>
                }
              />
            </FormControl>
            <br />
            <br />
            <div id="loginbtn">
              <Button variant="contained" color="primary" onClick={UserLogin}>
                登录
              </Button>
            </div>
          </div>
        </Fade>
      </Modal>
    </div>
  );
}
