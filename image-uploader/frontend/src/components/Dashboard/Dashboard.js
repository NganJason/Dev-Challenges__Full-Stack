import React, { useState } from "react"
import Upload from "../Upload/Upload";
import Uploading from "../Uploading/Uploading";
import Uploaded from "../Uploaded/Uploaded";
import ImgService from "../../model/imgService";

const imgService = new ImgService()

function Dashboard() {
  const [file, setFile] = useState(null)
  const [fileUrl, setFileUrl] = useState("")

  const onFileSubmit = async (file) => {
    setFile(file);

    let res = await imgService.uploadImg(file)
    setFile(null)
    console.log(res.data.url)
    setFileUrl(res.data.url);
  };

  const getComponent= () => {
    if (file === null && fileUrl === "") {
      return <Upload onFileSubmit={onFileSubmit} />;
    } else if (fileUrl !== "") {
      return <Uploaded url={fileUrl}/>
    }

    return <Uploading/>
  }

  return <div className="dashboard">
  {
    getComponent()  
  }
  </div>;
}

export default Dashboard;
