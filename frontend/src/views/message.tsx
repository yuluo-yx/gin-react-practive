import "./message.css";
import Button from "@material-ui/core/Button";
import { createStyles, Theme, makeStyles } from "@material-ui/core/styles";
import Input from "@material-ui/core/Input";
import InputLabel from "@material-ui/core/InputLabel";
import InputAdornment from "@material-ui/core/InputAdornment";
import FormControl from "@material-ui/core/FormControl";
import AccountCircle from "@material-ui/icons/AccountCircle";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      margin: theme.spacing(2),
      width: 200,
      fontSize: 30
    },
  })
);

export default function FullWidthGrid() {
  const classes = useStyles();
  return (
    <div>
      <div id="msg">msg</div>
      <div id="input">
        <FormControl className={classes.root}>
          <InputLabel className={classes.font}>
            Enter message information:{" "}
          </InputLabel>
          <Input
            id="input-with-icon-adornment"
            startAdornment={
              <InputAdornment position="start">
                <AccountCircle />
              </InputAdornment>
            }
          />
        </FormControl>
      </div>
    </div>
  );
}
