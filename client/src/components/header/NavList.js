import {
  Divider,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Avatar,
} from "@mui/material";

function NavList() {
  return (
    <>
      <div
        style={{ display: "flex", justifyContent: "center", padding: "10px" }}
      >
        <Avatar
          sx={{ bgcolor: "orange", width: "100px", height: "100px" }}
          alt='Remy Sharp'
          src='/logo.svg'
        >
          B
        </Avatar>
      </div>
      <Divider sx={{ backgroundColor: "white" }} />
      <List>
        {[
          "Minhas Notas",
          "Notas Públicas",
          "Minhas Comunidades",
          "Comunidades Populares",
          "Amigos",
        ].map((text) => (
          <ListItem key={text} disablePadding>
            <ListItemButton>
              <ListItemIcon></ListItemIcon>
              <ListItemText primary={text} />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </>
  );
}

export default NavList;
