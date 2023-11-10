import { Avatar as ChakraAvatar, AvatarProps } from "@chakra-ui/react";

function getRandomColor(str: string) {
  let asciiSum = 0;
  for (let i = 0; i < str.length; i++) {
    asciiSum += str.charCodeAt(i);
  }

  const red = Math.abs(Math.sin(asciiSum) * 256).toFixed(0);
  const green = Math.abs(Math.sin(asciiSum + 1) * 256).toFixed(0);
  const blue = Math.abs(Math.sin(asciiSum + 2) * 256).toFixed(0);
  return (alpha: number) => `rgba(${red}, ${green}, ${blue}, ${alpha})`;
}

const getGradientStyle: (text: string) => AvatarProps = (text) => {
  const color = getRandomColor(text);
  const color1 = color(1);
  const color2 = color(0.7);

  return {
    backgroundImage: `linear-gradient(135deg, ${color1}, ${color2})`,
    color: "white",
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    fontWeight: "bold",
    borderRadius: "10px",
    background: "transparent",
    textShadow: "1px 1px 3px rgba(0, 0, 0, 0.2)",
    fontFamily: "Helvetica, Arial, sans-serif",
    position: "relative",
    _before: {
      content: `""`,
      position: "absolute",
      left: "0",
      top: "0",
      w: "full",
      h: "full",
      borderRadius: "10px",
      backgroundColor: "white",
      backgroundImage: `repeating-radial-gradient( circle at 0 0, transparent 0, #ffffff 9px ), repeating-linear-gradient( ${color2}, ${color1} )`,
      zIndex: "-1",
    },
  };
};

const Avatar: React.FC<{ name: string } & AvatarProps> = ({
  name,
  ...rest
}) => {
  return <ChakraAvatar {...getGradientStyle(name)} {...rest} name={name} />;
};

export default Avatar;
