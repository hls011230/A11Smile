pragma solidity ^0.8.0;

contract UploadMedicalrecords{

//用户 
struct user_Users {
    address user;
    bool exit;
}

//上传医疗数据信息时间
uint public lastUpdateTime;

//医疗数据信息
struct user_Medicalinformation{
    address user;
    string[] PictureRoute;
}

//modified上传时间
modifier user_upTime(address _user){
    lastUpdateTime = block.timestamp;

    _;
}

//mapping对应用户
mapping (address => user_Users)  private Ausers;

//mapping对应医疗数据信息
mapping (address => user_Medicalinformation) private AmedicalInformation;

//设置用户
function user_Adduser(address _user) public{
       Ausers[_user].user= _user;
       Ausers[_user].exit= true;
}

//上传医疗信息
function user_AddMedicalInformation(address user,string memory Proute) public user_upTime(user) {
           require(user_IsUser(user),"patient no exist");
           AmedicalInformation[user].user=user;
           AmedicalInformation[user].PictureRoute.push(Proute);
}

//查询用户是否存在
function user_IsUser(address _user) public view returns(bool){
   return Ausers[_user].exit;
}

//查看病人上传的病历
function user_ViewMedicalRecords(address _user) public view returns(string[] memory){
    require(user_IsUser(_user),"patient no exist");
    return AmedicalInformation[_user].PictureRoute;
}

}