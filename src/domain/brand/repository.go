package brand

type BrandRepositoryIFace interface {
	Save(domain *Brand) (err error)
}
