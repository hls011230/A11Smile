pragma solidity ^0.8.0;



contract UploadMedicalrecords{



//用户 

struct Users {

  address user;

  bool exit;

}



//上传医疗数据信息时间

uint public lastUpdateTime;



//医疗数据信息

struct Medicalinformation{

  address user;

  string[] PictureRoute;

}



//modified上传时间

modifier upTime(address _user){

  lastUpdateTime = block.timestamp;



  _;

}



//mapping对应用户

mapping (address => Users)  private Ausers;



//mapping对应医疗数据信息

mapping (address => Medicalinformation) private AmedicalInformation;



//设置用户

function Adduser(address _user) public{

​    Ausers[_user].user= _user;

​    Ausers[_user].exit= true;

}



//上传医疗信息

function AddMedicalInformation(address user,string memory Proute) public upTime(user) {

​      require(IsUser(user),"patient no exist");

​      AmedicalInformation[user].user=user;

​      AmedicalInformation[user].PictureRoute.push(Proute);

}



//查询用户是否存在

function IsUser(address _user) public view returns(bool){

  return Ausers[_user].exit;

}



//查看病人上传的病历

function ViewMedicalRecords(address _user) public view returns(string[] memory){

  require(IsUser(_user),"patient no exist");

  return AmedicalInformation[_user].PictureRoute;

}



}