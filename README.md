# 微博广告 Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/biz-weibo.svg)](https://pkg.go.dev/github.com/bububa/biz-weibo)
[![Go](https://github.com/bububa/biz-weibo/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/biz-weibo/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/biz-weibo/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/biz-weibo/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/biz-weibo.svg)](https://github.com/bububa/biz-weibo)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/biz-weibo)](https://goreportcard.com/report/github.com/bububa/biz-weibo)
[![GitHub license](https://img.shields.io/github/license/bububa/biz-weibo.svg)](https://github.com/bububa/biz-weibo/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/biz-weibo.svg)](https://GitHub.com/bububa/biz-weibo/releases/)

- Oauth 授权 (api/oauth)
  - 生成授权链接 [ Authorize(clt *core.SDKClient, req *oauth.AuthorizeRequest) string ]
  - 获取 AccessToken [ AccessToken(clt *core.SDKClient, req *oauth.AccessTokenRequest) (*oauth.Token, error) ]
  - 刷新 Token [ RefreshToken(clt *core.SDKClient, req *oauth.RefreshTokenRequest) (*oauth.Token, error)]
- 账号服务 (api/account)
  - 获取广告主信息 [ Info(clt *core.SDKClient, accessToken string, advertiserID int64) (*account.Account, error) ]
  - 获取广告主账户流水信息 [ Flow(clt *core.SDK, accessToken string, req *account.FlowRequest) (*account.FlowsResponse, error) ]
  - 获取账户资产 [ Asset(clt *core.SDKClient, accessToken string, advertiserID int64) (*account.Asset, error) ]
  - 同步用户配置到服务器 [ SettingsSet(clt *core.SDKClient, accessToken string, req *account.SettingsSetRequest) (*account.Settings, error) ]
  - 获取用户配置 [ Settings(clt *core.SDKClient, accessToken string, req *account.SettingsRequest) (*account.Settings, error) ]
  - 资质列表 [ Licenses(clt *core.SDKClient, accessToken string, req *account.LicensesRequest) (*account.LicensesResponse, error) ]
- 广告投放
  - 账户日限额
    - 修改账户日限额 [ account.Budget(clt *core.SDKClient, accessToken string, req *account.BudgetRequest) (*account.Account, error) ]
    - 修改账户次日日限额 [ account.ReverseSpendCap(clt *core.SDKClient, accessToken string, req *account.ReverseSpendCapRequest) (*account.Account, error) ]
    - 删除账户次日日限额 [ account.DeleteReverseBudget(clt *core.SDKClient, accessToken string, advertiserID int64) (*account.Account, error) ]
  - 广告系列 (api/campaign)
    - 广告系列列表 [ Campaigns(clt *core.SDKClient, accessToken string, req *campaign.CampaignsRequest) (*campaign.CampaignsResponse, error) ]
    - 广告系列详情 [ Campaign(clt *core.SDKClient, accessToken string, req *campaign.CampaignRequest) (*campaign.Campaign, error) ]
    - 创建广告系列 [ Create(clt *core.SDKClient, accessToken string, req *campaign.CreateRequest) (*campaign.Campaign, error) ]
    - 更新系列预算或名称 [ Update(clt *core.SDKClient, accessToken string, req *campaign.UpdateRequest) (*campaign.Campaign, error) ]
    - 更新系列状态 [ UpdateStatus(clt *core.SDKClient, accessToken string, req *campaign.UpdateStatusRequest) ([]campaign.UpdateStatusResult, error) ]
  - 广告计划 (api/ad)
    - 广告计划列表 [ Ads(clt *core.SDKClient, accessToken string, req *ad.AdsRequest) (*ad.AdsResponse, error) ]
- 数据报表 (api/insights)
  - 效果统计分析 [ Effect(clt *core.SDKClient, accessToken string, req *insights.EffectRequest) (*insights.EffectResponse, error) ]
- 数据上报管理 (api/track)
  - 转化回传 [ Activate(clt *core.SDKClient, accessToken string, req *track.ActivateRequest) error ]
