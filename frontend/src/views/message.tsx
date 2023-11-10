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
import Msg from "../mock/messageData";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
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

const msgs: Msg[] = [
  {
    userId: "1",
    userName: "yuluo",
    message: "祝你平安-yuluo！",
  },
  {
    userId: "2",
    userName: "test",
    message: "祝你平安-test！",
  },
  {
    userId: "3",
    userName: "huakai",
    message: "祝你平安-huakai！",
  },
];

export default function FullWidthGrid() {
  const classes = useStyles();

  const listItem = msgs.map((item: Msg) => {
    return (
      <ListItem alignItems="flex-start">
        <ListItemAvatar>
          <Avatar className={classes.purple}>
            {getStr(0, 1, item.userName)}
          </Avatar>
        </ListItemAvatar>
        <ListItemText
          primary={item.userName}
          secondary={
            <React.Fragment>
              <Typography
                component="span"
                variant="body2"
                className={classes.inline}
                color="textPrimary"
              >
                {item.message}
              </Typography>
            </React.Fragment>
          }
        />
      </ListItem>
    );
  });

  return (
    <div>
      <h3>Message Board</h3>
      <div id="msg">{listItem}</div>
      <div id="input">
        <TextField
          className={classes.root}
          id="outlined-basic"
          label="Message"
          variant="outlined"
          defaultValue="祝你平安！"
          InputProps={{
            startAdornment: (
              <InputAdornment position="start">
                <Avatar className={classes.purple}>Y</Avatar>
              </InputAdornment>
            ),
          }}
        />
        <Button className={classes.btn} variant="contained" color="primary">
          发送
        </Button>
      </div>
    </div>
  );
}
