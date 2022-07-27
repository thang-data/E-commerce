import Repository from "./repository";

const repositories = [] as Array<Repository>;

type Constructor<T> = new (...args: any[]) => T;

const RepositoryFactory = {
    getRepository<R extends Repository>(type: Constructor<R>): R {
      const repository = repositories.find((a) => a instanceof type);
      if (repository) return repository as R;
      else {
        const newRepository = new type();
        repositories.push(newRepository);
        return newRepository;
      }
  
    },
  };
  
  export default RepositoryFactory;