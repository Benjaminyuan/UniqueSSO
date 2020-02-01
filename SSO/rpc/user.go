
package user 
import(
	"context"
	"google.golang.org/grpc"
	"unique/jedi/conf"
	pb "unique/jedi/protos/gencode"
	log "github.com/sirupsen/logrus"
)

const (
	addr = "localhost:50051"
)
type UserClient  struct {
	c pb.UserClient
}
func (client *UserClient)SayHello(ctx context.Context,req *pb.HelloRequest)(res *pb.HelloResponse,err error){
	res, err = client.c.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return nil,err
	}
	log.Infof("SayHello:%+v",res)
	return res,nil
}
func(client *UserClient)CreateUser(ctx context.Context,req *pb.CreateUserRequest)(res *pb.CreateUserResponse,err error){
	res,err = client.c.CreateUser(ctx,req)
	if err != nil {
		log.Fatalf("RPC CreateUser failed: %v", err)
		return nil,err
	}
	log.Infof("CreateUser:%+v",res)
	return res,nil
}
func (client *UserClient)GetUserInfo(ctx context.Context,req *pb.GetUserInfoByIDRequest)(res *pb.GetUserInfoByIDResponse,err error) {
	res,err = client.c.GetUserInfoByID(ctx,req)
	if err != nil {
		log.Fatalf("RPC GetUserInfoByID failed: %v", err)
		return nil,err
	}
	log.Infof("GetUserInfoByID :%+v",res)
	return res,nil
}
func (client *UserClient)VerifyUserIdentity(ctx context.Context,req *pb.VerifyUserIdentityRequest)(res *pb.VerifyUserIdentityResponse,err error) {
	res,err = client.c.VerifyUserIdentity(ctx,req)
	if err != nil {
		log.Fatalf("RPC VerifyUserIdentity failed: %v", err)
		return nil,err
	}
	log.Infof("VerifyUserIdentity :%+v",res)
	return res,nil
}
func NewUserClient()(*UserClient,error){
	log.Infof("addr:%v",conf.SSOConf.RPCConf.Addr)
	conn, err := grpc.Dial(conf.SSOConf.RPCConf.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil,err
	}
	return &UserClient{c:pb.NewUserClient(conn)},nil
}